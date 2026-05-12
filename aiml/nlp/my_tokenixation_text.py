import nltk
from nltk.tokenize import sent_tokenize
from nltk.tokenize import word_tokenize


corpus="""Hello Welcome,to kush naiks nlp token!
please do watch to become expert in nlp"""

#print(corpus)

documents=sent_tokenize(corpus)
type(documents)

# for sentence in documents:
#     print(sentence)

doc1=word_tokenize(corpus)
print(doc1)