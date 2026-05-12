'''Real-World Example:Multithreading for I/O-bound Tasks
Scenario:Web Scraping
Web Scraping often involves making numerous network requests to
fetch web pages.These tasks are I/O-bound because they spend a lot of time waiting for responses from servers.
Multithreading can significantly improve the. performance by allowing
multiple web pages to e fetched concurrently'''
'''
https://www.langchain.com/blog/documentation-refresh-for-langchain-v0-2
https://docs.langchain.com/langsmith/model-configurations
https://docs.langchain.com/langsmith/set-up-hierarchy
'''
import threading
import requests
from bs4 import BeautifulSoup

urls=[
    "https://www.langchain.com/blog/documentation-refresh-for-langchain-v0-2",
    "https://docs.langchain.com/langsmith/model-configurations",
    "https://docs.langchain.com/langsmith/set-up-hierarchy"
]

def fetch_content(url):
    response=requests.get(url)
    soup=BeautifulSoup(response.content,'html.parser')
    print(f'Fetched {(len(soup.text))} characters from {url}')

threads=[]

for url in urls:
    thread=threading.Thread(target=fetch_content,args=(url,))
    threads.append(thread)
    thread.start()

for thread in threads:
    thread.join()

print("All web pages fetched")