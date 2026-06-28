package ticket

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type ticketServiceTestEnv struct {
	tickets     *fakeTicketRepository
	flows       *fakeFlowRepository
	steps       *fakeFlowStepRepository
	agents      *fakeAgentRepository
	agentQueues *fakeAgentQueueRepository
}

func newTicketServiceTestEnv() (*TicketService, *ticketServiceTestEnv) {
	repos := &ticketServiceTestEnv{
		tickets: &fakeTicketRepository{
			tickets:       make(map[int64]*models.Ticket),
			hasInProgress: make(map[int64]bool),
			nextID:        1,
		},
		flows: &fakeFlowRepository{
			flows: map[int64]*models.Flow{
				1: {ID: 1, Name: "Default", IsActive: true},
			},
		},
		steps: &fakeFlowStepRepository{
			stepsByFlow: map[int64][]*models.FlowStep{
				1: {{ID: 1, FlowID: 1, QueueID: 10, Position: 1, Name: "First"}},
			},
			stepsByID: map[int64]*models.FlowStep{
				1: {ID: 1, FlowID: 1, QueueID: 10, Position: 1, Name: "First"},
			},
		},
		agents: &fakeAgentRepository{
			agents: map[int64]*models.Agent{
				1: {ID: 1, Name: "agent@example.com", IsAvailable: true},
			},
		},
		agentQueues: &fakeAgentQueueRepository{
			exists: map[[2]int64]bool{
				{1, 10}: true,
			},
		},
	}

	return NewTicketService(repos.tickets, repos.flows, repos.steps, repos.agents, repos.agentQueues), repos
}

type fakeTicketRepository struct {
	tickets       map[int64]*models.Ticket
	hasInProgress map[int64]bool
	nextID        int64
}

func (r *fakeTicketRepository) Create(_ context.Context, ticket *models.Ticket) error {
	ticket.ID = r.nextID
	r.nextID++
	r.tickets[ticket.ID] = cloneTicket(ticket)
	return nil
}

func (r *fakeTicketRepository) List(_ context.Context) ([]*models.Ticket, error) {
	tickets := make([]*models.Ticket, 0, len(r.tickets))
	for _, ticket := range r.tickets {
		tickets = append(tickets, cloneTicket(ticket))
	}
	return tickets, nil
}

func (r *fakeTicketRepository) FindByID(_ context.Context, id int64) (*models.Ticket, error) {
	ticket, ok := r.tickets[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	return cloneTicket(ticket), nil
}

func (r *fakeTicketRepository) UpdateState(_ context.Context, ticket *models.Ticket) error {
	if _, ok := r.tickets[ticket.ID]; !ok {
		return apperrors.ErrNotFound
	}
	r.tickets[ticket.ID] = cloneTicket(ticket)
	return nil
}

func (r *fakeTicketRepository) ListByStatus(_ context.Context, status string, _ int) ([]*models.Ticket, error) {
	tickets := make([]*models.Ticket, 0)
	for _, ticket := range r.tickets {
		if ticket.Status == status {
			tickets = append(tickets, cloneTicket(ticket))
		}
	}
	return tickets, nil
}

func (r *fakeTicketRepository) HasInProgressForAgent(_ context.Context, agentID int64) (bool, error) {
	return r.hasInProgress[agentID], nil
}

func (r *fakeTicketRepository) add(ticket *models.Ticket) *models.Ticket {
	if ticket.ID == 0 {
		ticket.ID = r.nextID
		r.nextID++
	}
	r.tickets[ticket.ID] = cloneTicket(ticket)
	return cloneTicket(ticket)
}

type fakeFlowRepository struct {
	flows map[int64]*models.Flow
}

func (r *fakeFlowRepository) Create(_ context.Context, _ *models.Flow) error {
	return nil
}

func (r *fakeFlowRepository) CreateWithSteps(_ context.Context, _ *models.Flow, _ []*models.FlowStep) error {
	return nil
}

func (r *fakeFlowRepository) Delete(_ context.Context, _ int64) error {
	return nil
}

func (r *fakeFlowRepository) List(_ context.Context) ([]*models.Flow, error) {
	return nil, nil
}

func (r *fakeFlowRepository) FindByID(_ context.Context, id int64) (*models.Flow, error) {
	flow, ok := r.flows[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *flow
	return &clone, nil
}

type fakeFlowStepRepository struct {
	stepsByFlow map[int64][]*models.FlowStep
	stepsByID   map[int64]*models.FlowStep
}

func (r *fakeFlowStepRepository) ListByFlowID(_ context.Context, flowID int64) ([]*models.FlowStep, error) {
	steps := r.stepsByFlow[flowID]
	result := make([]*models.FlowStep, len(steps))
	for i, step := range steps {
		clone := *step
		result[i] = &clone
	}
	return result, nil
}

func (r *fakeFlowStepRepository) FindByID(_ context.Context, id int64) (*models.FlowStep, error) {
	step, ok := r.stepsByID[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *step
	return &clone, nil
}

type fakeAgentRepository struct {
	agents map[int64]*models.Agent
}

func (r *fakeAgentRepository) List(_ context.Context) ([]*models.Agent, error) {
	return nil, nil
}

func (r *fakeAgentRepository) FindByID(_ context.Context, id int64) (*models.Agent, error) {
	agent, ok := r.agents[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *agent
	return &clone, nil
}

func (r *fakeAgentRepository) FindAvailableByQueueID(_ context.Context, _ int64) (*models.Agent, error) {
	return nil, nil
}

func (r *fakeAgentRepository) CreateForUser(_ context.Context, _ *models.User) error {
	return nil
}

type fakeAgentQueueRepository struct {
	exists map[[2]int64]bool
}

func (r *fakeAgentQueueRepository) Create(_ context.Context, _ *models.AgentQueue) error {
	return nil
}

func (r *fakeAgentQueueRepository) Delete(_ context.Context, _ *models.AgentQueue) error {
	return nil
}

func (r *fakeAgentQueueRepository) FindByAgentID(_ context.Context, _ int64) ([]*models.AgentQueue, error) {
	return nil, nil
}

func (r *fakeAgentQueueRepository) FindByQueueID(_ context.Context, _ int64) ([]*models.AgentQueue, error) {
	return nil, nil
}

func (r *fakeAgentQueueRepository) Exists(_ context.Context, agentID int64, queueID int64) (bool, error) {
	return r.exists[[2]int64{agentID, queueID}], nil
}

func cloneTicket(ticket *models.Ticket) *models.Ticket {
	clone := *ticket
	if ticket.CurrentFlowStepID != nil {
		value := *ticket.CurrentFlowStepID
		clone.CurrentFlowStepID = &value
	}
	if ticket.AssignedAgentID != nil {
		value := *ticket.AssignedAgentID
		clone.AssignedAgentID = &value
	}
	return &clone
}
