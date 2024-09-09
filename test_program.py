import requests
import random
import json

# Define possible types, recipients, and messages
types = ["email", "sms", "push"]
recipients = ["Saahil", "Alex", "Jordan", "Taylor", "Morgan"]
messages = [
    "Please do it.",
    "Don't forget to check the updates.",
    "Can you send me the report?",
    "Let's meet at 3 PM.",
    "Remember to submit your assignment."
]

# Define the number of iterations
num_iterations = 100

# Loop to send POST requests
for _ in range(num_iterations):
    # Randomly choose a type, recipient, and message
    chosen_type = random.choice(types)
    chosen_recipient = random.choice(recipients)
    chosen_message = random.choice(messages)

    # Create the payload
    payload = {
        "recipient": chosen_recipient,
        "message": chosen_message
    }

    # Define the URL
    url = f"http://localhost:8080/send/{chosen_type}"

    # Send the POST request
    try:
        response = requests.post(url, json=payload, verify=False)
        response.raise_for_status()  # Raise an error for bad responses
    except requests.exceptions.RequestException as e:
        print(f"Error for {chosen_type} to {chosen_recipient}: {e}")