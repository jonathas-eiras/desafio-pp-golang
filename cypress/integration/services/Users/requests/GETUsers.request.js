/// <reference types="cypress" />

const token = Cypress.env('token')

function allUsers() {
  return cy.request({
    method: 'GET',
    url: '/users',
    auth:{
          bearer: token
        } 
  })
}

export { allUsers };