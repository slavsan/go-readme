Feature: Generate markdown readme

  Scenario: Generates title
    Given a readme.yaml with contents
    """
    readme:
      title: My project title
    """
    When I parse the file
    Then the following README.md will be generated
    """
    # My project title
    """

  Scenario: Generates description
    Given a readme.yaml with contents
    """
    readme:
      description: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    """
    When I parse the file
    Then the following README.md will be generated
    """
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    """
