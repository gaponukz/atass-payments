from flask import Blueprint

test = Blueprint("test", __name__)


@test.get("/")
def test_controller():
    return "Hello, world!"
