const express = require('express')
const { graphqlHTTP } = require('express-graphql')
const { GraphQLSchema, GraphQLObjectType, GraphQLString } = require('graphql')

const app = express()

const schema = new GraphQLSchema({
    query: new GraphQLObjectType({
        name: 'Hello',
        fields: () => ({
            message: { 
                type: GraphQLString,
                resolve: () => 'Hello world'
            },
            name: { 
                type: GraphQLString,
                resolve: () => 'Your name'
            },
        })
    })
})

//For interactive graphql
app.use('/graphql', graphqlHTTP({
    schema: schema,
    graphiql: true
}))

app.listen(5000, () => console.log('Server running'))