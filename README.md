#Steganography with SendGrid Parse Webhook

This hack lets you encode the body of your email in the attachment. It requires steghide to be installed on your server.  

## How it works
 Using the SendGrid Parse Webhook, we can write code that gets attachments and parts of the email. In this case I take the body and encode it into the attachment using a unix program called Steghide. It uses a concept of security called Steganography. 
     If you attached a .JPG file, it will modify the pixels to store its data. Otherwise if you attached a .wav file, it will convert the text into audio samples and embed that into your file. Its really cool way of hiding information in plain sight. 


###  Get Set Up
- Install the SendGrid Go Library 
- Install StegHide on your server




