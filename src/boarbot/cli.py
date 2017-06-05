import argparse
import asyncio

from boarbot.client import BoarBotClient
from boarbot.dispatch import initialize_modules

def parse_cli() -> (str, ):
    parser = argparse.ArgumentParser()
    parser.add_argument('token', metavar='TOKEN', )
    args = parser.parse_args()
    return (args.token, )

def main():
    (token, ) = parse_cli()
    client = BoarBotClient()
    initialize_modules(client)
    client.run(token)

if __name__ == '__main__':
    main()
