def is_strong_password(password):
    if len(password)<=8:
        return False
    if not any(char.isdigit() for char in password):
        return Flase
    if not any(char.islower() for cahr in password):
        return Flase
    if not any(char.isupper()for cahr in password):
        return Flase
    if not any(char in '!@#$%^&*()+_-' for char in password):
        return False
    return True

print(is_strong_password("weakPwd"))