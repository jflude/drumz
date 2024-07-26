# Interview Exercise - Drum Machine

## The Objective
The goal of this exercise is to create a simple [drum machine sequencer](http://en.wikipedia.org/wiki/Drum_machine) that allows you to "play" your own version of the famous [four-on-the-floor](http://en.wikipedia.org/wiki/Four_on_the_floor_(music)) pattern, as shown below:

![https://raw.githubusercontent.com/mattetti/sm-808/master/Four_to_the_floor_Roland_TR-707.jpg](https://raw.githubusercontent.com/mattetti/sm-808/master/Four_to_the_floor_Roland_TR-707.jpg)

We'd like you to generate a real-time visual representation of a **hi-hat, bass drum and snare drum** playing the pattern. It's entirely up to you how to accept that input - you could use a text-based format, a binary one, or even accept interactive input from the user when your program starts. Different approaches will produce different trade-offs, so make sure you can discuss your thought process and implementation.

In terms of output, you'll need to display:

1. The name of the song being played
2. The tempo, represented as an integer value of "beats per minute"
3. The sequence of instruments, roughly in time with the tempo

Here's an example of a text-based implementation - notice that the instruments are being "played" at a consistent cadence, as opposed to the output just rendering all at once:

![https://raw.githubusercontent.com/mattetti/sm-808/master/drummachine-kata.gif](https://raw.githubusercontent.com/mattetti/sm-808/master/drummachine-kata.gif)

## Evaluation Criteria
We request that you use [the Go programming language](https://www.golang.org/) to implement your solution.

Once we receive your submission, we will anonymize it and then evaluate it using a scoring rubric. There are no tricks or "gotchas" in the evaluation - we'll be asking straightforward questions like:

1. Were we able to easily run your solution and see it working?
2. Does the solution handle unexpected inputs gracefully?
3. Has the code been tested and documented?
4. Are the key domain concepts modeled well?
5. Is the code written in a clear and consistent style?
6. Could we extend this solution in the future if requirements changed?

### Nice to Haves
Please note that **none** of the following items are required. These are just a few ideas that extend the exercise in interesting ways, at your discretion:

* Come up with additional songs and patterns to play. Can you replicate a part of your favorite song?
* A frontend interface could be interesting, possibly something browser-based for convenience.
* Output some real audio from the terminal! macOS ships with the `say` command, and Windows has `ptts`. There are also language bindings available for [portaudio](http://www.portaudio.com/).
* Being able to control the volume of an instrument is always helpful.

## Submitting Your Solution
Once you've reached a point where you feel your solution is complete, please compress all the relevant contents into an archive and upload it to the Greenhouse link provided. The hiring manager will remove potentially identifying information before submitting to engineering staff for scoring by rubric.