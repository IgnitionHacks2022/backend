import requests
import base64

with open("bottle.jpg", "rb") as img_file:
  #b64_string = base64.b64encode(img_file.read())

  url = "https://m.media-amazon.com/images/I/710w8zOFhLL._AC_SX355_.jpg"
  b64_string = base64.b64encode(requests.get(url).content)

  x = requests.post("http://localhost:8080/classify/abcdef", json = {"contents": b64_string})

  print(x.text)



