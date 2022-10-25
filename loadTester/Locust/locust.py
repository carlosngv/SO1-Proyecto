import json
from random import randrange
from locust import HttpUser, task, between

debug = False

def print_debug(msg):
    if debug:
        print(msg)

class Reader():
    def __init__(self) -> None:
        self.arr = []

    def pick_random(self):
        length = len(self.arr)
        if(length > 0):
            random_index = randrange(0, length - 1) if length > 1 else 0
            return self.arr.pop(random_index)
        else:
            print("Empty data...")
            return None

    def load(self):
        try:
            with open('../data.json', 'r') as data_file:
                self.arr = json.loads(data_file.read())
        except Exception as error:
            print(error)

class APIUser(HttpUser):

    wait_time = between(0.1, 0.9)
    reader = Reader()
    reader.load()

    @task
    def locust_page(self):
        random_data = self.reader.pick_random()
        if random_data is not None:
            data_to_send = json.dumps(random_data)
            print_debug(data_to_send)
            self.client.post(url="/locust_data", json=random_data)
        else:
            self.stop(True)
