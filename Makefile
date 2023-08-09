BRIEFCASE = briefcase

.PHONY: run setup

run:
	$(BRIEFCASE) dev

setup:
	python -m venv venv
	source venv/bin/activate
	pip install -r requirements.txt
