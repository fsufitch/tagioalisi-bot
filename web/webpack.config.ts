// Built/comprehended with thanks to:
// * Webpack Basics: https://medium.com/age-of-awareness/setup-react-with-webpack-and-babel-5114a14a47e9#a8f2
// * Why use Babel? https://stackoverflow.com/a/49624611
// * Babel+Typescript explainer: https://github.com/Microsoft/TypeScript-Babel-Starter

import 'webpack-dev-server';
import * as path from 'path';
import { Configuration } from 'webpack';
import { fileURLToPath } from 'url';

import HtmlWebpackPlugin from 'html-webpack-plugin';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const configureTypescript = async (prod: boolean): Promise<Configuration> => {
    const tsLoader = { loader: 'ts-loader' };  // https://www.npmjs.com/package/ts-loader
    const babelLoader = {
        // https://www.npmjs.com/package/babel-loader
        loader: 'babel-loader',
        options: {
            cacheDirectory: true,
            presets: [
                ['@babel/preset-env', { useBuiltIns: 'entry', corejs: 3 }],
                '@babel/preset-react',
            ],
            plugins: [],
        },
    };
    const sourceMapLoader = { loader: 'source-map-loader' }

    const module = {
        rules: [{
            test: /\.tsx?$/i, use: [
                babelLoader,
                ...(prod ? [] : [sourceMapLoader]),
                tsLoader,
            ]
        }],
    }

    const { default: TsconfigPathsPlugin } = await import('tsconfig-paths-webpack-plugin');
    const resolve = {
        // Add `.ts` and `.tsx` as a resolvable extension.
        extensions: [".ts", ".tsx", ".js", ".jsx"],
        plugins: [new TsconfigPathsPlugin()],
    }

    const entry = {
        'app': path.join(__dirname, 'src', 'app.tsx'),
    }

    // const { default: ForkTsCheckerWebpackPlugin } = await import('fork-ts-checker-webpack-plugin');

    // const plugins = [
    //     new ForkTsCheckerWebpackPlugin()
    // ]

    return { entry, resolve, module };
}

const configureStyles = async (prod: boolean): Promise<Configuration> => {
    let sassLoader = { loader: 'sass-loader', options: { sourceMap: true, implementation: 'sass' } };
    let cssLoader = { loader: 'css-loader', options: { sourceMap: true, modules: true } };
    let postCssLoader = { loader: 'postcss-loader', options: { sourceMap: true } };
    let resolveUrlLoader = { loader: 'resolve-url-loader', options: { sourceMap: true } };

    const { default: MiniCssExtractPlugin } = await import('mini-css-extract-plugin');
    const miniCssExtractPlugin = new MiniCssExtractPlugin({
        filename: '[name].css',
    });
    const styleLoader = 'style-loader'; // Cannot import it as it's missing type declarations

    const { default: CssMinimizerPlugin } = await import('css-minimizer-webpack-plugin');
    const cssMinimizerPlugin = new CssMinimizerPlugin();

    // Using style-loader in development makes HMR work better
    const actualStyleLoader = prod ? MiniCssExtractPlugin.loader : styleLoader;
    const stylePlugins = prod ? [miniCssExtractPlugin] : [];

    return {
        module: {
            rules: [
                { test: /\.s?css$/i, use: [actualStyleLoader, cssLoader, postCssLoader, resolveUrlLoader, sassLoader] },
            ]
        },
        resolve: {
            alias: {
                'tagioalisi': path.join(__dirname, 'src', 'tagioalisi'),
            }
        },
        plugins: stylePlugins,
        optimization: {
            minimizer: ['...', cssMinimizerPlugin],
        }
    };
}


const configureHTML = async (prod: boolean): Promise<Configuration> => {
    const htmlWebpackPlugin = new HtmlWebpackPlugin({
        template: path.join(__dirname, "src", "index.html"),
        xhtml: true,
        minify: prod,
    })

    return { plugins: [htmlWebpackPlugin] }
}

const configureAssets = async (prod: boolean): Promise<Configuration> => ({
    module: {
        rules: [
            { test: /\.png/i, type: 'asset/resource' },
            { test: /\.svg/i, type: 'asset/resource' },
        ]
    },
});


const configureBaseWebpack = async (prod: boolean): Promise<Configuration> => {
    const { BundleAnalyzerPlugin } = await import('webpack-bundle-analyzer');  // https://github.com/webpack-contrib/webpack-bundle-analyzer
    return ({
        mode: prod ? 'production' : 'development',
        devtool: prod ? false : 'eval-cheap-module-source-map', // https://webpack.js.org/configuration/devtool/
        output: {
            path: path.join(__dirname, 'dist'),
            filename: "[name].bundle.js",
        },
        optimization: {
            minimize: prod,
            splitChunks: {
                chunks: "all",
                name: 'vendor',
            },
        },
        plugins: [
            ...(!!process.env.WEBPACK_ANALYZER ? [new BundleAnalyzerPlugin()] : []),
        ],
        devServer: {
            hot: false,  // React is not setup for this
            liveReload: true,
            port: process.env.WEBPACK_DEV_PORT || 8080,
            historyApiFallback: true,
            open: false,
            headers: {
                'Set-Cookie': `BOT_EXTERNAL_BASE_URL=${process.env.BOT_EXTERNAL_BASE_URL}; SameSite=Lax`,
            },
        },
    })
};

// ===========

type ConfigurationFunction = (prod: boolean) => Promise<Configuration>;

const buildConfig = async (env: any, argv: { [key: string]: any }, funcs: Iterable<ConfigurationFunction>) => {
    const prod = argv.mode == 'production';

    // Set the env var as well, to inform Babel
    process.env.NODE_ENV = prod ? 'production' : 'development';

    const { merge } = await import('webpack-merge');

    let config: Configuration = {};
    for (const func of funcs) {
        config = merge(config, await func(prod))
    }

    return config;
}

export default async (env: any, argv: any) => {
    const config = await buildConfig(env, argv, [
        configureBaseWebpack,
        configureTypescript,
        configureHTML,
        configureStyles,
        configureAssets,
    ])
    // console.log(config);
    return config;
}