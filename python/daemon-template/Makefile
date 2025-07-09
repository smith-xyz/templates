.PHONY: help install install-dev test lint format type-check clean run

help:
	@echo "Available commands:"
	@echo "  install     - Install production dependencies"
	@echo "  install-dev - Install development dependencies"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linter"
	@echo "  format      - Format code"
	@echo "  type-check  - Run type checker"
	@echo "  clean       - Clean build artifacts"
	@echo "  run         - Run the application"

install:
	pip install -r requirements.txt

install-dev:
	pip install -r requirements.txt
	pip install -e .[dev]

test:
	pytest

lint:
	flake8 .

format:
	black .

type-check:
	mypy .

clean:
	rm -rf build/
	rm -rf dist/
	rm -rf *.egg-info/
	find . -type d -name __pycache__ -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete

run:
	python main.py 