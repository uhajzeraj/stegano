Steganography
=======

# IMT2681 Assignment 3

[![Build Status](https://travis-ci.org/uhajzeraj/stegano.svg?branch=master)](https://travis-ci.org/uhajzeraj/stegano)

# Steganography Services and Other Classic Cryptographic Algorithms

Our project idea was to create a web service where users could use some of the most famous classical encryption techniques.
The main service we offer is Steganography. The other services offered are Caesar's cipher & ROT13 algorithm.

### For this project we have used:

  * Heroku
  * OpenStack
  * MongoDB

### During our work in this project we learned a lot of new things: 

  * We learned how to better use Go language as a back-end programming language, implementing together with it all of the front-end components such as HTML, Javascript, CSS, Bootstrap & jQuery.
  * How to implement basic cryptographic algorithms into GO lang, such as Steganography and some substitution algorithms.
  * How to work with Mongo Databases and how to save the users information in them.
  * Implementing new things in OpenStack, which are explained in more details later in this Readme.
  * Working together as a group, organizing the time correctly and the work between members of our group.
  
## Members of our group and the total work hours dedicated to the project cumulatively by the group
  
  Uran Hajzeraj (uhajzeraj)     - 22 hours
  Nehar Jashari (neharj)        - 18 hours
  Besnik Kelmendi (besnikk)     - 19 hours
  Gresa Dodaj (gresad)          - 18 hours
  Etnik Gashi (etnikg)          - 19 hours
  
Total of 96 hours of work for the group.


## Heroku deployment

The app has been deployed in Heroku and has this link: https://imt2681-stegano.herokuapp.com/


# Usage of the services

First you have to create an account so that the necessary information about the users are saved into our database. The passwords of the users are saved in our database using the most advanced hashing techniques **bcrypt**. **bcrypt** is a password hashing function designed by Niels Provos and David Mazières, based on the Blowfish cipher, and presented at USENIX in 1999. Besides incorporating a salt to protect against rainbow table attacks, bcrypt is an adaptive function: over time, the iteration count can be increased to make it slower, so it remains resistant to brute-force search attacks even with increasing computation power. Read more about it in [here](https://en.wikipedia.org/wiki/Bcrypt?fbclid=IwAR02_QdFVS8AgzDLpw4SsRgvqec-gww7aoj2t01bsfh1slKuNIf5LF0Oi2c).

Screenshot 1        -       Screenshot 2

After creating an account you can access the services offered by us.

  Screenshot 3

## Steganography

Steganography is the hiding of a secret message within an ordinary message and the extraction of it at its destination. Basically is the practice of concealing a file, message, image, or video within another file, message, image, or video. Steganography takes cryptography a step farther by hiding an encrypted message so that no one suspects it exists. Ideally, anyone scanning your data will fail to know it contains encrypted data.

You can read more about Stegaography in this link: <a href="https://en.wikipedia.org/wiki/Steganography">Steganography Wikipedia</a>

In our website you can enter the icon for steganography and follow these steps to use it:

1. Select an image which you want to hide text to.
  Screenshot 7
  
2. Choose an option. Do you want to encode text into that image or if you know there is text hidden with steganography you want to decode it.
  Screenshot 8
  
3. If you choose Encoding, then you can write any text you want in that image and that text will never be seen by the ordinary human eye. If you want to Decode that image then the hidden text will be shown to the user after pressing that button.
  Screenshot 9
  
4. After encoding the image, that image will be saved in the particular user database along with his other information. The saved data will be accessible for the user in the Saved Data page in the website. You can download that image for further use or you can delete it.
  Screenshot 10


## Caesar's Cipher

In cryptography, a Caesar cipher is one of the simplest and most widely known encryption techniques. It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet.

The method is named after Julius Caesar, who used it in his private correspondence.

In our site you can encrypt text using the Caesar's Cipher, while choosing the shifting size for that or you can decrypt a text with this algorithm if you know the correct shifting size used for the ecryption.

  Screenshot 5
  
## ROT 13

ROT13 ("rotate by 13 places", sometimes hyphenated ROT-13) is a simple letter substitution cipher that replaces a letter with the 13th letter after it, in the alphabet.

ROT13 is a special case of the Caesar cipher which was developed in ancient Rome.

In our site you can encrypt text using this algorithm with pressing the Encrypt text button shown in the webpage.

  Screenshot 6

## User page

In this page you can see the information that is saved by us in our database and you can do other actions concerning your account.
You can change current password and you can delete your account.

  Screenshot 4


# Resources

[Official MongoDB driver](https://github.com/mongodb/mongo-go-driver)
[MGO](https://github.com/globalsign/mgo)
[bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
[Gorilla Mux](https://github.com/gorilla/mux)


# Build with

**Back end**:   Go Language

**Front end**:  HTML, CSS, Javascript, Bootstrap & jQuery.
