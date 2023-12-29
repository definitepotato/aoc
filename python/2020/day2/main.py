#!/usr/bin/env python3

with open('input.txt') as input:
  entry = input.readlines()
  entry = [entry.strip() for entry in entry]


class Password:
  def __init__(self, policy_and_password):
    self.policy = policy_and_password.split(':')[0]
    self.password = policy_and_password.split(':')[1]
    self.policy_letter = self.policy.split(' ')[1]
    self.policy_range = self.policy.split(' ')[0]
    self.policy_range_low = int(self.policy_range.split('-')[0])
    self.policy_range_high = int(self.policy_range.split('-')[1])

  def count_letters(self) -> int:
    count = 0
    for letter in self.password:
      if letter == self.policy_letter:
        count += 1
    return count

  def is_valid(self) -> bool:
    if self.count_letters() >= self.policy_range_low and self.count_letters() <= self.policy_range_high:
      return True
    else:
      return False
  
  def is_valid_position(self) -> bool:
    valid = False
    if self.password[self.policy_range_low] == self.policy_letter or self.password[self.policy_range_high] == self.policy_letter:
      valid = True
    if self.password[self.policy_range_low] == self.policy_letter and self.password[self.policy_range_high] == self.policy_letter:
      valid = False
    return valid

count = 0
for policy_and_password in entry:
  p = Password(policy_and_password)

  if p.is_valid_position():
    count += 1

print(count)

# def count_letters(policy_letter, password) -> int:
#   count = 0
#   for letter in password:
#     if letter == policy_letter:
#       count += 1
#   return count


# for entry in e:
#   policy = entry.split(':')[0]
#   password = entry.split(': ')[1]

#   policy_letter = policy.split(' ')[1]
#   policy_range = policy.split(' ')[0]
#   policy_range_low = int(policy_range.split('-')[0])
#   policy_range_high = int(policy_range.split('-')[1])

#   number_of_letters = count_letters(policy_letter, password)

#   if number_of_letters >= policy_range_low and number_of_letters <= policy_range_high:
#     print(f'{policy}: {password} = valid')
#   else:
#     print(f'{policy}: {password} = invalid')