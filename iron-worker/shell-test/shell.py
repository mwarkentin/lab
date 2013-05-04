import json
from time import asctime, gmtime
from uuid import uuid4

import envoy
from github import Github
from jinja2 import Environment, FunctionLoader
import requests

from iron_utils import get_task_id, get_payload


def get_branch_name(repo):
    """
    Generates a branch name from repo count stat.

    If the stats service is unavailable, generate the branch name with
    a UUID.

    """
    url = "{domain}/repo/{repo}/count".format(
        domain=payload['stats_domain'],
        repo=repo,
    )
    branch_format = 'shrinkray-{}'

    try:
        r = requests.post(url)
        count_json = r.json()
        branch_name = branch_format.format(count_json['count'])
    except (requests.ConnectionError, KeyError):
        branch_name = branch_format.format(uuid4())
        print "exception=requests.ConnectionError function=get_branch_name branch_name={}".format(
            branch_name,
        )

    return branch_name


def get_changed_images(lines):
    images = []
    lines = [line for line in lines if 'bytes' in line]

    for image in lines:
        # [1:] removes initial . from image name
        name = image.split(' ')[0][1:]
        if 'unchanged' in image:
            print 'Unable to shrink image: {0}'.format(image)
        else:
            size = parse_metrics(image)
            images.append({'name': name, 'size': size})
    return images


def parse_metrics(image):
    numbers = [int(s) for s in image.split() if s.isdigit()]
    try:
        reduced, remaining = numbers
    except ValueError:
        return None
    original = remaining + reduced
    return {'original': original, 'reduced': reduced, 'remaining': remaining}


def record_metrics(repo, images):
    task_id = get_task_id()
    print "Recording metrics for {0}\n".format(repo)
    data = json.dumps({'images': images})
    print data

    url = "{domain}/repo/{repo}".format(
        domain=payload['stats_domain'],
        repo=repo,
    )

    headers = {'content-type': 'application/json'}
    try:
        r = requests.post(url, data=data, headers=headers)
        print "POST task={} url={} status_code={}".format(
            task_id,
            url,
            r.status_code,
        )
    except requests.ConnectionError:
        print "exception=requests.ConnectionError task={} function=record_metrics url={}".format(
            task_id,
            url,
        )


def create_pull_request(oauth_key, repo, branch_name, images):
    body = generate_body(images)
    g = Github(oauth_key)
    r = g.get_repo(repo)
    r.create_pull('Shrinkray', body, r.master_branch, branch_name)


def generate_body(images):
    template = env.get_template('body.md')
    return template.render(images=images)


def get_template(name):
    return """
## Optimized images

{% for image in images %}
* {{ image.name }} ({{ image.size.original }}b -> {{ image.size.remaining }}b)
{% endfor %}
"""


env = Environment(loader=FunctionLoader(get_template))
payload = get_payload()
print payload

split_repo = payload['repo'].split('/')

branch_name = get_branch_name(payload['repo'])
print "branch_name={}".format(branch_name)

cmd = './test.sh {} {} {} {}'.format(
    split_repo[0],
    split_repo[1],
    payload['oauth_key'],
    branch_name
)
print cmd

r = envoy.run(cmd)
lines = r.std_out.splitlines()
images = get_changed_images(lines)
print lines
print images
if images:
    record_metrics(payload['repo'], images)
    create_pull_request(payload['oauth_key'], payload['repo'], branch_name, images)
else:
    print "No images could be shrunk"
print "Shrink all images end: {0}".format(asctime(gmtime()))
