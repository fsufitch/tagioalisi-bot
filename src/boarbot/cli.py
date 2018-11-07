import argparse
import asyncio

from boarbot.bootstrap.bootstrap import run_bootstrap
from boarbot.db.common import BoarbotDatabase
from boarbot.client import BoarBotClient
from boarbot.dispatch import initialize_modules

def parse_cli() -> (str, ):
    parser = argparse.ArgumentParser()
    parser.add_argument('token', metavar='TOKEN', help='Discord API token identifying your bot')
    parser.add_argument('-m', '--module', dest='modules', metavar='import.path:ClassName',
                        action='append', default=[],
                        help='import paths and class names of modules to load; can be specified multiple times')
    args = parser.parse_args()
    return (args.token, args.modules)

def main():
    (token, modules) = parse_cli()
    client = BoarBotClient()
    db = BoarbotDatabase.get_instance()
    run_bootstrap(db.get_session())
    initialize_modules(client, modules)
    client.run(token)

if __name__ == '__main__':
    main()
