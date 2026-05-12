import streamlit as st
import pandas as pd

st.title("Streamlit Text Input")

name=st.text_input("Enter yout name: ")

age=st.slider("Select your age: ",0,100,18)


options=["Python","Java","C++","JavaScript"]
choice=st.selectbox("Choose your favraite laungugae: ",options)


if name:
    st.write(f"hello {name}")
st.write(f"You Selected {choice}.")
st.write(f"Your age is {age}.")


data={
    "Name":["John","Jane","Jake","Jill"],
    "Age": [28,24,35,40],
    "City":["New York","Los Angeles","Chicago","Houston"]
}

df=pd.DataFrame(data)
df.to_csv("sampledata.csv")
st.write(df)

uploaded_file=st.file_uploader("Choose a CSV file",type="csv")
if uploaded_file is not None:
    df=pd.read_csv(uploaded_file)
    st.write(df)