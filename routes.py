from flask import Blueprint, render_template, redirect, url_for, request, flash
from models import User, db

main = Blueprint('main', __name__)

@main.route('/')
def index():
    return "Welcome to PortfoAI"

@main.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        pass
    return render_template('login.html')
