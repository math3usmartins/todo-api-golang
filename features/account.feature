Feature: User account
  In order to manage to-do items
  An account must be created

  @account
  Scenario: Sign up with success
    When the account service is called to create a user with the following information:

    | email         | password |
    | new@email.com | abc123   |

    Then a user can be found by email "new@email.com"
