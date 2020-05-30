import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="Barq", # Replace with your own username
    version="0.0.1",
    author="Manuel Lorente",
    author_email="manloralm@outlook.com",
    description="See long description",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/manulorente/barq",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.8',
)