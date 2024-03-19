MailerService
=============

# Intro
This service is used for sending out mails when connected to a SMTP server in the backend

This service should expose and endpoint which can be requested for sending an Email.
This service requires an existing SMTP credential for the backend.
Sending mail as a service should be handled by a celery worker.
