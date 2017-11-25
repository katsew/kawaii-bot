FROM python:3-alpine

WORKDIR /usr/src/app

COPY ./heartcatch/ ./
#RUN pip install --no-cache-dir -r requirements.txt
RUN pip install -r requirements.txt

CMD ["python", "app.py"]
