import streamlit as st
import pandas as pd
import numpy as np

##Display header of the application
st.title("Streamlit Text Input ")

## Text Input
name=st.text_input("Enter Your name: ")

if name:
    st.write(f"Hello, {name}!")

##slider
age=st.slider("Select your age: " ,0,100,25)

st.write(f"Your age is: {age}")

option=["java","python","c++","javascript"]
choice=st.selectbox("Select your favorite programming language: ",option)
st.write(f"Your favorite programming language is: {choice}")

data={
    'Name':['Alice','Bob','Charlie','David'],
    'Age':[25,30,35,40],
    'City':['New York','Los Angeles','Chicago','Houston']
}

df=pd.DataFrame(data)
st.write("DataFrame:")
df.to_csv("Sampledata.csv")
st.write(df)

uploaded_file=st.file_uploader("Upload a CSV file", type=["csv"])
if uploaded_file is not None:
    df=pd.read_csv(uploaded_file)
    st.write(df)