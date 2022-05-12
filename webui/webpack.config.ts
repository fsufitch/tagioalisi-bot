// Built/comprehended with thanks to:
// * Webpack Basics: https://medium.com/age-of-awareness/setup-react-with-webpack-and-babel-5114a14a47e9#a8f2
// * Why use Babel? https://stackoverflow.com/a/49624611
// * Babel+Typescript explainer: https://github.com/Microsoft/TypeScript-Babel-Starter

import * as path from 'path';
import { Configuration, optimize } from 'webpack';

const configureTypescript = async (prod: boolean): Promise<Configuration> => {
    const tsLoader = { loader: 'ts-loader' };  // https://www.npmjs.com/package/ts-loader
    const babelLoader = {
        // https://www.npmjs.com/package/babel-loader
        loader: 'babel-loader',
        options: {
            cacheDirectory: true,
            presets: [
                ['@babel/preset-env', { useBuiltIns: 'entry', corejs: 3 }],
                '@babel/preset-react'
            ],
        }
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


    return { entry, resolve, module };
}

const configureStyles = async (prod: boolean): Promise<Configuration> => {
    let sassLoader = { loader: 'sass-loader', options: { sourceMap: true } };
    let cssLoader = { loader: 'css-loader', options: { sourceMap: true, modules: true } };
    let postCssLoader = { loader: 'postcss-loader', options: { sourceMap: true } };
    let resolveUrlLoader = { loader: 'resolve-url-loader', options: {sourceMap: true}};

    const { default: MiniCssExtractPlugin, loader: miniCssExtractLoader } = await import('mini-css-extract-plugin');
    const miniCssExtractPlugin = new MiniCssExtractPlugin({
        filename: '[name].css',
    });

    const { default: CssMinimizerPlugin } = await import('css-minimizer-webpack-plugin');
    const cssMinimizerPlugin = new CssMinimizerPlugin();

    return {
        module: {
            rules: [
                { test: /\.scss$/i, use: ['style-loader', cssLoader, postCssLoader, resolveUrlLoader, sassLoader] },
            ]
        },
        resolve: {
            alias: {
                'tagioalisi-styles': path.join(__dirname, 'src', 'tagioalisi-styles'),
            }
        },
        plugins: [miniCssExtractPlugin],
        optimization: { minimizer: ['...', cssMinimizerPlugin] }
    };
}


const configureHTML = async (prod: boolean): Promise<Configuration> => {
    const HtmlWebpackPlugin = await import('html-webpack-plugin');
    const htmlWebpackPlugin = new HtmlWebpackPlugin.default({
        template: path.join(__dirname, "src", "index.html"),
        xhtml: true,
        minify: prod,
    })

    return { plugins: [htmlWebpackPlugin] }
}

const configureAssets = async (prod: boolean): Promise<Configuration> => ({
    module: {
        rules: [
            {
                test: /\.png/,
                type: 'asset/resource'
            }
        ]
    },
});

const configureBaseWebpack = async (prod: boolean): Promise<Configuration> => ({
    mode: prod ? 'production' : 'development',
    devtool: prod ? 'source-map' : 'source-map', // ???
    output: {
        path: path.join(__dirname, 'dist'),
        filename: "[name].bundle.js",
    },
    optimization: {
        splitChunks: {
            chunks: "all",
            maxSize: 50000,
            name: 'vendor',
        },
        minimize: true,
    }
});

// ===========

type ConfigurationFunction = (prod: boolean) => Promise<Configuration>;

const buildConfig = async (env: any, argv: { [key: string]: any }, funcs: Iterable<ConfigurationFunction>) => {
    const prod = argv.mode == 'production';
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
    console.log(config);
    return config;
}