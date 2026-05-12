from flask import Flask, render_template, request
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
    return render_template('index.html')

@app.route("/about")
def about():
    return render_template('about.html')

@app.route("/form", methods=['GET', 'POST'])
def form():
    if request.method == 'POST':
        name = request.form['name']
        return f'Hello, {name}!'
    return render_template('form.html')
    # return '''
    #     <form method="post">
    #         Name: <input type="text" name="name">
    #         <input type="submit" value="Submit">
    #     </form>
    # '''

if __name__ == '__main__':
    app.run(debug=True)