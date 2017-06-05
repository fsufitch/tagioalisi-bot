from setuptools import setup, Extension, find_packages

setup(name='boarbot',
      version='1.0',
      author='Blackshell',
      author_email="fsufitchi@gmail.com",
      description="Custom bot for the Sociologic Planning Boar",
      url="https://github.com/fsufitch/discord-boar-bot",
      package_dir={'': 'src'},
      packages=find_packages('src'),
      package_data={'': ['*']},
      entry_points={
          "console_scripts": [
              "boarbot=boarbot.cli:main"
          ],
        },

      install_requires=[
          'discord.py',
          ],
      )
