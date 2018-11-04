# Alkalarm-alexa-skills

This is the alexa skill for the alkalarm system project integration.

The main idea of this project is create an integration using alexa skills and the echo device
to control the "alkalarm" home security system.

For example, we could manage the alarm system using:

_**Alexa, open the alarm system, and activate it, after 30 second**_

_**Alexa, open the alarm system, and activate it just for the perimeter**_

_**Alexa, open the alarm system, and stop it**_

_**Alexa, tell me the alarm system state**_





![Alarm System ](./images/schema.jpg)

## Alexa Skills Voice Processing Architecture

Just to keep in mind the steps that we have to do in order to create and integrate custom skills with the alkalarm
project, we're gonna review the main architecture of alexa skills processing:

![Alexa Architecture](https://cdn-images-1.medium.com/max/1600/1*2K8S9Zjh2ZQVyRE3gBHK-A.jpeg)

As you can see, we have to define 3 things:

1-. Create the skills definition in AWS Alexa development site.

2-. Create and upload the lambda code to response the skills questios.

3-. Integrate your service (alkalarm) into the lambda processing step.



