import requests
import base64

with open("bottle.jpg", "rb") as img_file:
  b64_string = base64.b64encode(img_file.read())

  x = requests.post("http://localhost:8080/classify/abcdef", json = {"contents": b64_string})

  print(x.text)



