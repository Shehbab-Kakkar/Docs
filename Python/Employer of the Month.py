#Program to get the Employer of the Month
employee = [("Ram",200),("Mohan",100),("John",50)]
def get_employee_of_month(employee):
      WorkingHours = 0
      empOfMonth =  ''
      for emp,hours in employee:
        if hours > WorkingHours:
          empOfMonth = emp
          WorkingHours = hours
          
        else:
          pass
      
      return (empOfMonth,WorkingHours)        
print(get_employee_of_month(employee))       
name,prize = get_employee_of_month(employee)
print(name)
print(prize)

"""
('Ram', 200)
Ram
200
"""
