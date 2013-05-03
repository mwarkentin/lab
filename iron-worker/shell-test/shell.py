from uuid import uuid4

import envoy
import requests

from iron_utils import get_payload


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
print r.std_out
