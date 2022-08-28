import requests
import base64


url = "http://localhost:8080/register"

payload = {
  "email": "zhehaizhang3@gmail.com",
  "password": "zhehaizhang",
  "bluetoothid": "zhehaizhang",
  "name": "zhehai"
}

res = requests.post(url, json=payload)

url = "http://localhost:8080/login"

payload = {
  "email": "zhehaizhang3@gmail.com",
  "password": "zhehaizhang",
}

res = requests.post(url, json=payload)




resJson = res.json()

token = resJson["token"]

headers = {
  'content-type': 'application/json',
  'token': token
}

x = requests.post("http://localhost:8080/statistics", json = {}, headers=headers)

print(x.text)



