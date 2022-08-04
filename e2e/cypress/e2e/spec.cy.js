describe('empty spec', () => {
  it('Visits my app', () => {
    cy.visit('http://frontend:8080')
    

  cy.get('input').first()
      .type('Task by cypress')
      .should('have.value', 'Task by cypress')
  cy.get('button').first().click()
  cy.get('input').first().clear() 
  cy.get('input').first()
      .type('Task 2')
    

  cy.get('button').first().click()
  cy.get('li').should('contain', 'Task by cypress')
  cy.get('[type="checkbox"]').last().check()

  cy.get('div').get('button').last().click()
  
  cy.get('div').should('not.contain', 'Task 2')
   
   
  })
})