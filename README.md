# fizzbuzz

This fizzbuzz uses concurrency to create a "high throughput" fizzbuzz
application.

Here is a simplified diagram of what's happening:

                       +--> fizzer ----+
                       |               |
    for first..last ---+--> buzzer ----+--> results
                       |               |
                       +--> numberer --+

For other fizzbuzzes, see [this](https://play.golang.org/p/JGW_6_5lMMi) and [that](https://play.golang.org/p/oMtKHXNfrZi).

And then there's [the reddit post](https://www.reddit.com/r/golang/comments/ixedgq/variations_on_fizzbuzz/).
