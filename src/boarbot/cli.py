import argparse

def parse_cli() -> (str, ):
    parser = argparse.ArgumentParser()
    parser.add_argument('token', metavar='TOKEN', )
    args = parser.parse_args()
    return (args.token, )

def main() -> None:
    print('hello!')
    (token, ) = parse_cli()
    print('got token', token)

if __name__ == '__main__':
    main()
