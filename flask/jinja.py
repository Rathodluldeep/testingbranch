## Building URL DYnamically
## variable rules
## Jinja 2 template engine

###Jinha 2 template engine is a powerful tool for rendering dynamic content in web applications. It allows you to create HTML templates with placeholders that can be filled with data at runtime. Jinja2 is the default template engine for Flask, a popular web framework in Python.
'''
{{ }}expressions to print output in html
{% %} statements for control flow, such as loops and conditionals
{# #} comments that are not rendered in the final output
'''


from jinja2 import Template
from flask import Flask, render_template, request, redirect,url_for
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

# @app.route("/submit", methods=['GET', 'POST'])
# def submit():
#     if request.method == 'POST':
#         name = request.form['name']
#         return f'Hello, {name}!'
#     return render_template('form.html')
    # return '''
    #     <form method="post">
    #         Name: <input type="text" name="name">
    #         <input type="submit" value="Submit">
    #     </form>
    # '''
##Variable Rules
@app.route('/success/<int:score>')
def success(score):
    res=""
    if score >= 50:
        res="You have passed the exam."
    else:
        res="You have failed the exam."
        
    return render_template('result.html', result=res)

##Variable Rules
@app.route('/successres/<int:score>')
def successres(score):
    res=""
    if score >= 50:
        res="You have passed the exam."
    else:
        res="You have failed the exam."
    exp = {'score': score, 'res': res}
        
    return render_template('result1.html', result=exp)
## If condition
@app.route('/successif/<int:score>')
def successif(score):
        
    return render_template('result.html', result=score)

@app.route('/fail/<int:score>')
def fail(score):
    return render_template('result.html', result=score)

@app.route('/submit', methods=['GET', 'POST'])
def submit():
    if request.method == 'POST':
        science = float(request.form['Science'])
        math = float(request.form['Math'])
        c = float(request.form['C'])
        datascience = float(request.form['datascience'])

        total = science + math + c + datascience
        average = total / 4

        return redirect(url_for('successres', score=average))

    return render_template('getresults.html')

if __name__ == '__main__':
    app.run(debug=True)
