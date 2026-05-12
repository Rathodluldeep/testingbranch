import streamlit as st
import pandas as pd

st.title("stremalit Text Input")

name=st.text_input("Enter your name:")

age=st.slider("select your age:" ,0,100,18)

st.write(f"your age is {age}.")

option=["Python","Java","C++","Javascript"]
choice=st.selectbox("Choose your favorite laungugae: ",option)
st.write(f"You selected {choice}.")
if name:
    st.write(f"Hello,{name}")

data={
    "Name": ["John","Jane","Jack","Jill"],
    "Age" : [28,25,35,40],
    "City": ["New York","Los Angles","Chicago","Houston"]
}

df=pd.DataFrame(data)
st.write(df)

Uploaded_file=st.file_uploader("CHoose a CSV file",type="csv")
if Uploaded_file is not None:
    df=pd.read_csv(Uploaded_file)
    st.write(df)