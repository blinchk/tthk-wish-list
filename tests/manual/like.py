import json
import requests
import bcolors
import os

for i in range(100):
    headers = {
        "Token": os.getenv("TOKEN"),
        "Content-Type": "application/json"
    }
    wish = {
        "id": 7,
    }
    payload = {
        "connection": wish["id"],
        "connection_type": "wishes"
    }
    url = "http://localhost:8080"
    endpoint = "/wish/like"
    try:
        r = requests.post(url + endpoint,
                          data=json.dumps(payload),
                          headers=headers,
                          timeout=1.2)
    except requests.ReadTimeout:
        print(f"{bcolors.ERRMSG}Timed out{bcolors.END} Request #{i}")
        continue
    print(f"{bcolors.OKMSG}SUCCESS{bcolors.END} Request #{i}\nTime: {r.elapsed.total_seconds()}s\nBody: {r.text}")
