from flask import Flask
'''
It create an instance of the Flask class,
which is the WSGI (web server gateway interface) application.
The first argument is the name of the application's module or package.'''

##WSGI application
app = Flask(__name__)

@app.route("/")
def welcome():
    return 'Welcome to the Flask App!'

@app.route("/index")
def index():
    return 'Welcome to the index page!'

if __name__ == '__main__':
    app.run(debug=True)