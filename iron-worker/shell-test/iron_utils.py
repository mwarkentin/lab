import json
import sys


def get_payload():
    payload_file = None
    payload = None

    for i in range(len(sys.argv)):
        if sys.argv[i] == "-payload" and (i + 1) < len(sys.argv):
            payload_file = sys.argv[i + 1]
            with open(payload_file, 'r') as f:
                payload = json.loads(f.read())
            break
    return payload


def get_task_id():
    """
    Get the IronWorker task id
    """

    for i in range(len(sys.argv)):
        if sys.argv[i] == "-id" and (i + 1) < len(sys.argv):
            return sys.argv[i + 1]
