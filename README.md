# fizzbuzz

This fizzbuzz uses concurrency to create a "high throughput" fizzbuzz
application.

Here is a simplified diagram of what's happening:

                       +--> fizzer ----+
                       |               |
    for first..last ---+--> buzzer ----+--> results
                       |               |
                       +--> numberer --+
