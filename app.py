from flask import Flask
from src.infrastructure.controllers import test

app = Flask(__name__)
app.register_blueprint(test.test)

if __name__ == "__main__":
    app.run()
