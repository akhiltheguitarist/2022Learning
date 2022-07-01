const express = require('express')
const { graphqlHTTP } = require('express-graphql')
const { GraphQLSchema, GraphQLObjectType, GraphQLString, GraphQLList, GraphQLInt, GraphQLNonNull } = require('graphql')

const app = express()

const authors = [{id:1, name: 'Sam'}, {id:2, name: 'Tom'}, {id:3, name: 'Kane'}]
const books = [
    {id:1, name:'Do great things', authorId: 1},
    {id:2, name:'My world', authorId: 1},
    {id:3, name:'The earth', authorId: 2},
    {id:4, name:'Sports Diary', authorId: 3},
    {id:5, name:'Moneyball', authorId: 3}
]

const AuthorType = new GraphQLObjectType(
    {
        name: 'Author',
        description: 'Author representation',
        fields: () => ({
            id: { type: GraphQLNonNull(GraphQLInt)},
            name: { type: GraphQLNonNull(GraphQLString)},
        })
    }
)

const BookType = new GraphQLObjectType(
    {
        name: 'book',
        description: 'Book representation',
        fields: () => ({
            id: { type: GraphQLNonNull(GraphQLInt)},
            name: { type: GraphQLNonNull(GraphQLString)},
            authorId: { type: GraphQLNonNull(GraphQLInt)},
            author: { type: AuthorType, resolve: (book) => { return authors.find(author => author.id===book.authorId)}}
        })
    }
)

const RootQueryType = new GraphQLObjectType({
    name: 'query',
    description: 'Root Query',
    fields: () => ({
        books: {
            type: new GraphQLList(BookType),
            description: 'List of books',
            resolve: () => books
        }
    })
})

const schema = new GraphQLSchema({
    query: RootQueryType
})

//For interactive graphql
app.use('/graphql', graphqlHTTP({
    schema: schema,
    graphiql: true
}))

app.listen(5000, () => console.log('Server running'))