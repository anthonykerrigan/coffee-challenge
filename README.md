# Coffee challenge

This task it to build an app to fetch and parse some data associated with a coffee shop.

The app keeps track of coffee ordered, what the balance is for each user, what users have paid for already, and what is still owed.

A command line tool that simply logs output to the terminal is sufficient, however, you are welcome to build the app in whatever way you see fit.  You are welcome to use whatever language, packages, and frameworks that you like.

## Data

We've got the following data
- `data/prices.json` - provided by our barista. Has details of what beverages are available, and what their prices are.
- `data/orders.json` - list of beverages ordered by users of the app.
- `data/payments.json` - list of payments made by users paying for items they have purchased.

## Requirements

- Load the list of prices

- Load the orders

- Calculate the total cost of each user's orders

- Load the payments

- Calculate the total payment for each user

- Calculate what each user now owes

- Log the results to the command line (or display them in some UI if you would rather build something with a UI).  These results should take the form of an array of objects with the keys of: `user`, `order_total`, `payment_total`, and `balance`, for example:

  ```
  [
    { "user": "coach",    "order_total": 5.00, "payment_total": 2.50, "balance": 2.50 },
    { "user": "ellis",    "order_total": 6.15, "payment_total": 6.15, "balance": 0.00 },
    { "user": "rochelle", "order_total": 6.90, "payment_total": 4.50, "balance": 2.40 },
    { "user": "zoey",     "order_total": 2.30, "payment_total": 0.00, "balance": 2.30 }
  ]
  ```

## Submitting The Test

1. Ensure that everything you wish to submit is committed to the `master` branch
1. In your project directory, run: `tar -czvf firstname_lastname.tar.gz` (where _firstname_ and _lastname_ are your names)
1. Email the generated `.tar.gz` file back to the person that sent you the test.

## What We Care About

Here's what you should aim for with your code:
- Clean, readable, production quality code; would we want to work with your code as part of a bigger codebase?
- Solid testing approach
- Use Git
- Commit small changes often so we can see your approach, and progress.
- Include the `.git` directory in the packaged .tar.gz file you send to us.
- Include a README.md explaining how to run your code after it's extracted.
- We haven't hidden any nasty tricks in the test. Don't overthink it. Just write nice, solid code.
