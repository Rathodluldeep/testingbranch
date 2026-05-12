## PUT and DELETE methods for API
## working with API's --json
from flask import Flask, request, jsonify

app = Flask(__name__)

###INitital data in my todo list
items=[
    {'id':1, 'task':'Buy groceries'},
    {'id':2, 'task':'Clean the house'},
    {'id':3, 'task':'Finish the project'}
]

@app.route('/')
def home():
    return "Welcome to the To-Do List API!"


##Get:Retrieve all the items
@app.route('/items', methods=['GET'])
def get_items():
    return jsonify(items)

##Get:Retrieve a specific item by ID
@app.route('/items/<int:item_id>', methods=['GET'])
def get_item(item_id):
    item = next((item for item in items if item['id'] == item_id), None)
    if item:
        return jsonify(item)
    else:
        return jsonify({'message': 'Item not found'}), 404
    
##POST: Create a new item
@app.route('/items', methods=['POST'])
def create_item():
    if not request.json or not 'name'in request.json:
        return jsonify({'message': 'Bad request'}), 400
    new_item = {
        'id': items[-1]['id'] + 1 if items else 1,
        'task': request.json['name'],
        'description': request.json['description']
    }
    items.append(new_item)
    return jsonify(new_item), 201

##PUT: Update an existing item
@app.route('/items/<int:item_id>', methods=['PUT'])
def update_item(item_id):
    item = next((item for item in items if item['id'] == item_id), None)
    if not item:
        return jsonify({'message': 'Item not found'}), 404
    if not request.json:
        return jsonify({'message': 'Bad request'}), 400
    item['task'] = request.json.get('name', item['task'])
    item['description'] = request.json.get('description', item.get('description'))
    return jsonify(item)

##DELETE: Remove an existing item
@app.route('/items/<int:item_id>', methods=['DELETE'])
def delete_item(item_id):
    global items
    items = [item for item in items if item['id'] != item_id]
    return jsonify({'message': 'Item deleted'})

if __name__=='__main__':
    app.run(debug=True)