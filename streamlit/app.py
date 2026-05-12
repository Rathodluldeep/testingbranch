import streamlit as st
import pandas as pd
import numpy as np

##title of the application
st.title("hello Streamlit")

## Display a Simple Text
st.write("This is a simple text")

## Display a DataFrame
df=pd.DataFrame({
    'Column 1':[1,2,3,4],
    'Column 2':[10,20,30,40]
})

st.write(df)

## Display a Line Chart
charts_data=pd.DataFrame(np.random.randn(20,3),columns=['a','b','c'])
st.line_chart(charts_data)
