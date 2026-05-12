from flask import Flask,render_template
'''
It create an instance of the Flask class,
which is the WSGI (web server gateway interface) application.
The first argument is the name of the application's module or package.'''

##WSGI application
app = Flask(__name__)

@app.route("/")
def welcome():
    return '<html><h1>Welcome to the Flask App!</h1></html>'

@app.route("/index")
def index():
    return render_template('index.html')

@app.route("/about")
def about():
    return render_template('about.html')

if __name__ == '__main__':
    app.run(debug=True)