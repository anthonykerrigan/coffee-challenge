import json

def read_json_file(file_path):
    with open(file_path, 'r') as file:
        try:
            json_data = json.load(file)
            return json_data
        except json.JSONDecodeError:
            print(f"Error: Invalid JSON format in file: {file_path}")
            return None

def display_objects(json_data):
    if json_data is not None:
        for obj in json_data:
            print(obj)

pricesFile = './data/prices.json'  # Import the Prices to a variable
prices = read_json_file(pricesFile)

ordersFile = './data/orders.json'  # Import the Orders to a variable. 
orders = read_json_file(ordersFile)

# Display objects from the first JSON file
print("Objects from the first JSON file:")
display_objects(prices)

# Display objects from the second JSON file
print("Objects from the second JSON file:")
display_objects(orders)
