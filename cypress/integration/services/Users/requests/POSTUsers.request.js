/// <reference types="cypress" />
const faker = require('faker');
const payloadAddUser = require('../payloads/new-user.json')
const token = Cypress.env('token')

payloadAddUser.email = faker.internet.email()

function addUser() {
  return cy.request({
    method: 'POST',
    url: '/users',
    body: payloadAddUser,
    auth:{
            bearer: token
        } 
  })
}

export { addUser };