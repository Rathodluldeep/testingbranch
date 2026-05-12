class Youtube:
    def __init__(self,username,subscription=0):
        self.username=username
        self.subscription=subscription

user1=Youtube("Elshad",0)
# print(user1.username)
# print(user1.subscription)

user2=Youtube("Renad")
user2.subscription = 5
# print(user2.username)
# print(user2.subscription)

print(user1.__dict__)
print(user2.__dict__)
print(Youtube.__dict__)