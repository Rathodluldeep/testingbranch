#class StarCookies:
#     def __init__(self,color,weight):
#         #initialize attribute
#         self.color=color
#         self.weight=weight
        
        

# star_cookie1=StarCookies()
# star_cookie1.weight=15
# star_cookie1.color="Red"
# print(star_cookie1.weight)
# print(star_cookie1.color)

# star_cookie2=StarCookies()
# star_cookie2.weight=16
# star_cookie2.color="Blue"
# print(star_cookie2.weight)
# print(star_cookie2.color)

class Youtube:
    def __init__(self,username,subscription=0,subscriber=0):
        self.username=username
        self.subscription=subscription
        self.subscriber=subscriber
    def subscriber(self,user):
        user.subscriber +=1
        self.subscription +=1
        
        

user1=Youtube("Elshad",0)
# print(user1.username)
# print(user1.subscription)

user2=Youtube("Renad")
user2.subscription = 5
# print(user2.username)
# print(user2.subscription)
user1.subscription(user2)
user2.subscription(user1)

print(user1.__dict__)
print(user2.__dict__)
print(Youtube.__dict__)
