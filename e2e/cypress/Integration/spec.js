it('Visits my app', () => {
    cy.visit('')
    

  cy.get('input').first()
      .type('Task by cypress')
      .should('have.value', 'Task by cypress')

  cy.get('button').first().click()
  cy.get('li').should('contain', 'Task by cypress')
//   cy.get('[type="checkbox"]').last().check()

//   cy.get('li').get('button').last().click()
//   cy.get('li').should('not.contain', 'Task by cypress')
   
   
  })