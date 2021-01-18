##### 1. Import pandas and read in the Ecommerce Purchases csv file and set it to a DataFrame called ecom. 

In [1]:

```python
import pandas as pd
```

In [2]:

```python
ecom = pd.read_csv("Ecommerce Purchases.csv")
```



##### 2. Check the head of the DataFrame

In [3]:

```
ecom.head()
```



#### 3.  How many rows and columns are there? 

In [5]:

```python
ecom.info()
```



#### 4. What is the average Purchase Price? 

In [6]:

```python
ecom["Purchase Price"].mean()
```



#### 5. What were the highest and lowest purchase prices? 

In [7]:

```python
ecom["Purchase Price"].max()
```

In [8]:

```python
ecom["Purchase Price"].min()
```



#### 6. How many people have English 'en' as their Language of choice on the website? 

In [9]:

```python
ecom[ecom['Language'] == 'en'].count()
```



#### 7. How many people have the job title of "Lawyer" ? 

In [10]:

```python
ecom[ecom['Job'] == 'Lawyer'].info()
```



#### 8. How many people made the purchase during the AM and how many people made the purchase during PM ? 

**(Hint: Check out [value_counts()](http://pandas.pydata.org/pandas-docs/stable/generated/pandas.Series.value_counts.html) ) **

In [11]:

```python
ecom["AM or PM"].value_counts()
```



#### 9. What are the 5 most common Job Titles? 

In [12]:

```python
ecom["Job"].value_counts().head(5)
```



#### 10. Someone made a purchase that came from Lot: "90 WT" , what was the Purchase Price for this transaction? 

In [13]:

```python
ecom[ecom['Lot'] == "90 WT"]["Purchase Price"]
```



#### 11. What is the email of the person with the following Credit Card Number: '4926535242672853 '

In [14]:

```python
ecom[ecom["Credit Card"] == 4926535242672853]['Email']
```



#### 12. How many people have American Express as their Credit Card Provider \*and* made a purchase above $95 ?

In [15]:

```python
ecom[(ecom["CC Provider"] == "American Express") & (ecom["Purchase Price"] > 95)].count()
```

Out[15]:

```
Address             39
Lot                 39
AM or PM            39
Browser Info        39
Company             39
Credit Card         39
CC Exp Date         39
CC Security Code    39
CC Provider         39
Email               39
Job                 39
IP Address          39
Language            39
Purchase Price      39
dtype: int64
```



#### 13. Hard: How many people have a credit card that expires in 2025?

In [16]:

```python
sum(ecom["CC Exp Date"].apply(lambda x : int(x.split('/')[1]) == 25))
```

Out[16]:

```
1033
```



#### 14. Hard: What are the top 5 most popular email providers/hosts (e.g. gmail.com, yahoo.com, etc...) 

In [17]:

```python
ecom["Email"].apply(lambda x : x.split('@')[1]).value_counts().head(5)
```

Out[17]:

```
hotmail.com     1638
yahoo.com       1616
gmail.com       1605
smith.com         42
williams.com      37
Name: Email, dtype: int64
```