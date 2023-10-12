package k8sallowedrepos

violation[{"msg": msg}] {
    container := input.review.object.spec.containers[_]

    # Assume the image is not from an allowed repo
    image_is_allowed := false

    # Check if the container's image starts with any of the allowed repos
    repo = input.parameters.repos[_]
    startswith(container.image, repo)
    image_is_allowed := true
    
    # Condition for the violation is that image is not allowed
    not image_is_allowed

    # Generate the violation message
    msg := sprintf("container <%v> has an invalid image repo <%v>, allowed repos are %v", [container.name, container.image, input.parameters.repos])
}