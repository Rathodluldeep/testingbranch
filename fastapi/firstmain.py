from fastapi import FastAPI
from enum import Enum

app=FastAPI()

@app.get('/')
def index():
    return 'hello world!'

@app.get('/blog/all')
def get_all_blogs(page=1,page_size: Optional[int]= None):
    return {'message': f'All {page_size} blogs on pages{page}'}

class BlogType(str, Enum):
    short= 'short'
    story= 'story'
    howto= 'howto'

@app.get('/blog/type/{type}')
def get_blog_type(type: BlogType):
    return {'message': f'Blog Type {type}'}

@app.get('/blog/{id}')
def get_blog(id):
    return {"message": f"Blog with id {id}"}