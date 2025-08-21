# Original input string to be masked
input_str = "fe80::f0d0:d896:3214:e4d9%39"

# Initialize an empty string to store the masked result
masked_str = ''

# Loop through each character in the original string
for ch in input_str:
    # Check if the character is an alphabet letter (a-z or A-Z)
    if ('a' <= ch <= 'z') or ('A' <= ch <= 'Z'):
        masked_str += 'X'  # Replace letters with 'X'
    else:
        masked_str += ch  # Keep non-letters unchanged

# Print the final masked string
print(masked_str)
