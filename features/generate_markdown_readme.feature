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

  Scenario: Generates license
    Given a readme.yaml with contents
    """
    readme:
      license: MIT
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### License
    MIT
    """

  Scenario: Generates prerequisites
    Given a readme.yaml with contents
    """
    readme:
      prerequisites: |
        You will need to make sure
        you have the following prerequisites
        in order to install and use the project.
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Prerequisites
    You will need to make sure
    you have the following prerequisites
    in order to install and use the project.
    """

  Scenario: Generates author
    Given a readme.yaml with contents
    """
    readme:
      author: Slavcho Slavchev
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Author
    Slavcho Slavchev
    """

  Scenario: Generates install instructions
    Given a readme.yaml with contents
    """
    readme:
      install: |
        In order to install do the following
        ```
        npm install
        ```
        Then do this and that
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Installation
    In order to install do the following
    ```
    npm install
    ```
    Then do this and that
    """

  Scenario: Generates usage instructions
    Given a readme.yaml with contents
    """
    readme:
      usage: |
        In order to generate a README.md file from a readme.yaml, run
        ```
        go-readme create
        ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Usage
    In order to generate a README.md file from a readme.yaml, run
    ```
    go-readme create
    ```
    """

  Scenario: Generates clone info
    Given a readme.yaml with contents
    """
    readme:
      setup:
        clone: |
          ```
          git clone git@github.com:slavsan/go-readme.git
          ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Setup

    #### Clone
    ```
    git clone git@github.com:slavsan/go-readme.git
    ```
    """

  Scenario: Generates build info
    Given a readme.yaml with contents
    """
    readme:
      setup:
        build: |
          ```
          go build main.go
          ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Setup

    #### Build
    ```
    go build main.go
    ```
    """

  Scenario: Generates start info
    Given a readme.yaml with contents
    """
    readme:
      setup:
        start: |
          ```
          go run main.go
          ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Setup

    #### Start
    ```
    go run main.go
    ```
    """

  Scenario: Generates testing info
    Given a readme.yaml with contents
    """
    readme:
      setup:
        testing: |
          ```
          godog
          ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Setup

    #### Testing
    ```
    godog
    ```
    """

  Scenario: Generates lint info
    Given a readme.yaml with contents
    """
    readme:
      setup:
        linting: |
          ```
          golangci-lint run
          ```
    """
    When I parse the file
    Then the following README.md will be generated
    """
    ### Setup

    #### Linting
    ```
    golangci-lint run
    ```
    """
